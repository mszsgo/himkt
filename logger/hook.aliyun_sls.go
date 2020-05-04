package logger

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
	"time"
)

type aliyunSlsHook struct {
	AliyunSls *AliyunSls
}

func (hook *aliyunSlsHook) Fire(entry *logrus.Entry) error {
	bytes, err := entry.Bytes()
	if err != nil {
		return err
	}
	// 转换为Map，存入SLS
	var mkv map[string]string
	err = json.Unmarshal(bytes, &mkv)
	if err != nil {
		return err
	}
	hook.AliyunSls.SendLog(mkv)
	return nil
}

func (hook *aliyunSlsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

// 获取本机网卡IP
func getLocalIP() (ipv4 string) {
	var (
		err     error
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}

	if err != nil {
		ipv4 = "127.0.0.1"
	}
	return
}

type AliyunSls struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Project         string `json:"project"`
	Logstore        string `json:"logstore"`

	producer *producer.Producer `json:"-"`
	topic    string             `json:"-"`
	source   string             `json:"-"`
}

func (sls *AliyunSls) SetTopic(topic string) *AliyunSls {
	sls.topic = topic
	return sls
}

func (sls *AliyunSls) Producer() *producer.Producer {
	if sls.producer != nil {
		return sls.producer
	}
	sls.source = getLocalIP()

	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = sls.Endpoint
	producerConfig.AccessKeyID = sls.AccessKeyID
	producerConfig.AccessKeySecret = sls.AccessKeySecret
	producer := producer.InitProducer(producerConfig)
	producer.Start() // 启动producer实例
	sls.producer = producer
	return sls.producer
}

func (sls *AliyunSls) SendLog(kv map[string]string) {
	log := producer.GenerateLog(uint32(time.Now().Unix()), kv)
	// himkt   msd   127.0.0.1   k-v
	err := sls.Producer().SendLog(sls.Project, sls.Logstore, sls.topic, sls.source, log)
	if err != nil {
		fmt.Println(err)
	}
}

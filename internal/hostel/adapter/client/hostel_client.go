package client

import (
	"context"
	"net/http"

	"hostel-service/internal/hostel/domain"

	"github.com/core-go/core/client"
)

type HostelClient struct {
	Client *http.Client
	Url    string
	Config *client.LogConfig
	Log    func(context.Context, string, map[string]interface{})
}

type ResultInfo struct {
	Status  int64          `mapstructure:"status" json:"status" gorm:"column:status" bson:"status" dynamodbav:"status" firestore:"status"`
	Errors  []ErrorMessage `mapstructure:"errors" json:"errors,omitempty" gorm:"column:errors" bson:"errors,omitempty" dynamodbav:"errors,omitempty" firestore:"errors,omitempty"`
	Message string         `mapstructure:"message" json:"message,omitempty" gorm:"column:message" bson:"message,omitempty" dynamodbav:"message,omitempty" firestore:"message,omitempty"`
}

type ErrorMessage struct {
	Field   string `mapstructure:"field" json:"field,omitempty" gorm:"column:field" bson:"field,omitempty" dynamodbav:"field,omitempty" firestore:"field,omitempty"`
	Code    string `mapstructure:"code" json:"code,omitempty" gorm:"column:code" bson:"code,omitempty" dynamodbav:"code,omitempty" firestore:"code,omitempty"`
	Param   string `mapstructure:"param" json:"param,omitempty" gorm:"column:param" bson:"param,omitempty" dynamodbav:"param,omitempty" firestore:"param,omitempty"`
	Message string `mapstructure:"message" json:"message,omitempty" gorm:"column:message" bson:"message,omitempty" dynamodbav:"message,omitempty" firestore:"message,omitempty"`
}

func NewHostelClient(config client.ClientConfig, log func(context.Context, string, map[string]interface{})) (*HostelClient, error) {
	c, _, conf, err := client.InitializeClient(config)
	if err != nil {
		return nil, err
	}
	return &HostelClient{Client: c, Url: config.Endpoint.Url, Config: conf, Log: log}, nil
}

func (c *HostelClient) Load(ctx context.Context) ([]domain.Hostel, error) {
	url := c.Url
	var hostels []domain.Hostel
	err := client.Get(ctx, c.Client, url, &hostels, c.Config, c.Log)
	return hostels, err
}

func (c *HostelClient) LoadByCode(ctx context.Context, code string) (*domain.Hostel, error) {
	url := c.Url + "/" + code
	var hostel domain.Hostel
	err := client.Get(ctx, c.Client, url, &hostel, c.Config, c.Log)
	return &hostel, err
}

func (c *HostelClient) Create(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	var res ResultInfo
	err := client.Post(ctx, c.Client, c.Url, hostel, &res, c.Config, c.Log)
	return res.Status, err
}

func (c *HostelClient) Update(ctx context.Context, hostel *domain.Hostel) (int64, error) {
	url := c.Url + "/" + hostel.Id
	var res ResultInfo
	err := client.Put(ctx, c.Client, url, hostel, &res, c.Config, c.Log)
	return res.Status, err
}

func (c *HostelClient) Delete(ctx context.Context, code string) (int64, error) {
	url := c.Url + "/" + code
	var res int64
	err := client.Delete(ctx, c.Client, url, &res, c.Config, c.Log)
	return res, err
}

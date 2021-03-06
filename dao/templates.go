package dao

import (
	"errors"
	"github.com/lxn/go-pgsql"
)

const (
	QUERY_LOAD_TEMPLATE_BY_ID = "select body from delivery.templates where template_id = @id;"
)

type Templates struct {
	*BaseDao
}

func (self *Templates) LoadByLang(lang string) (string, error) {
	return "{{.Body}}{{.UnsubscribeLink}}", nil
}

func (self *Templates) LoadTemplate(templateId int64) (string, error) {
	connection, err := self.getConnection()
	if err != nil {
		return "", err
	}
	var result string
	parameter := pgsql.NewParameter("@id", pgsql.Bigint)
	err = parameter.SetValue(templateId)
	recordSet, err := connection.Query(QUERY_LOAD_TEMPLATE_BY_ID, parameter)
	if err != nil {
		return "", err
	}
	defer recordSet.Close()
	fetched, err := recordSet.FetchNext()
	if err != nil {
		return "", err
	}
	if !fetched {
		return "", errors.New("Template not found")
	}

	err = recordSet.Scan(&result)
	if err != nil {
		return "", err
	}

	return result, nil
}

func NewTemplatesDao(connectionString string) *Templates {
	return &Templates{&BaseDao{connectionString}}
}

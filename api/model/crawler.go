package model

type Crawler struct {
	Url      string //Site URI
	IsApi    bool   //True or False
	ApiParam string
	Regex    string //Regex finder in html page

	Value interface{} //Crawler value
	Type  string      //string or bool or int
}

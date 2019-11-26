/*
 *	convert String ID to primitive.ObjectID 
 *	same can be use to other objectID's as well
 *	forcus on mongodb queries  
 */

newsSources := [] string {"5dca3be9c5eaf07e9cce7b11","5dca3be9c5eaf07e9cce7b14"} 

var newsSourcesIds []primitive.ObjectID
var source string
for i := range newsSources {
	source = newsSources[i]
	res, errSpurceId := primitive.ObjectIDFromHex(source)
	if errSpurceId != nil {
		logger.Println("String to primitive object id Error")
		var parseObjectIdErrResponse newsfeedErrorResponse.ErrorResponse
		parseObjectIdErrResponse.Error = "Article List: Sources Id: " + source + ", errMsg: " + errSpurceId.Error()
		parseObjectIdErrResponse.Type = "ERR_PARSING_VALUE"
		return c.JSON(404, parseObjectIdErrResponse)
	}
	newsSourcesIds = append(newsSourcesIds, res)
}
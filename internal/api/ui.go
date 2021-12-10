package api

import (
	"fmt"
	"strings"

	"github.com/dollarkillerx/creeper/internal/conf"
	"github.com/gin-gonic/gin"
)

func (a *ApiServer) webUi(ctx *gin.Context) {

	index := a.ser.AllIndex()

	htm := strings.ReplaceAll(uiTemp, `{{$.Token}}`, conf.CONFIG.Token)

	idx := ""
	for i, v := range index {
		idx += fmt.Sprintf(`<option value="%d">%s</option>`, i, v)
	}

	htm = strings.ReplaceAll(htm, `{{$.Index}}`, idx)

	ctx.Header("content-type", "text/html charset=utf-8")
	ctx.Writer.Write([]byte(htm))
}

var uiTemp = `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>Creeper 简易数据分析</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- CSS only -->
    <!-- CSS only -->
    <script src="https://cdn.jsdelivr.net/npm/jquery@3.6.0/dist/jquery.min.js"></script>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/css/bootstrap.min.css" rel="stylesheet"
          integrity="sha384-F3w7mX95PdgyTmZZMECAngseQB83DfGTowi0iMjiWaeVhAn4FJkqJByhZMI3AhiU" crossorigin="anonymous">
    <!-- JavaScript Bundle with Popper -->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/js/bootstrap.bundle.min.js"
            integrity="sha384-/bQdsTh/da6pkI1MST/rWKFNjaCP5gBSY4sEBT38Q/9RBh9AH40zEOg7Hlq2THRZ"
            crossorigin="anonymous"></script>

    <style>
        .bd-placeholder-img {
            font-size: 1.125rem;
            text-anchor: middle;
            -webkit-user-select: none;
            -moz-user-select: none;
            user-select: none;
        }

        @media (min-width: 768px) {
            .bd-placeholder-img-lg {
                font-size: 3.5rem;
            }
        }
    </style>

</head>
<body>
<div class="container py-4">
    <input type="hidden" id="token" value="{{$.Token}}">

    <header class="pb-3 mb-4 border-bottom">
        <a href="#" class="d-flex align-items-center text-dark text-decoration-none">
            <svg xmlns="http://www.w3.org/2000/svg" width="40" height="32" class="me-2" viewBox="0 0 118 94" role="img">
                <title>Creeper 简易数据分析</title>
                <path fill-rule="evenodd" clip-rule="evenodd"
                      d="M24.509 0c-6.733 0-11.715 5.893-11.492 12.284.214 6.14-.064 14.092-2.066 20.577C8.943 39.365 5.547 43.485 0 44.014v5.972c5.547.529 8.943 4.649 10.951 11.153 2.002 6.485 2.28 14.437 2.066 20.577C12.794 88.106 17.776 94 24.51 94H93.5c6.733 0 11.714-5.893 11.491-12.284-.214-6.14.064-14.092 2.066-20.577 2.009-6.504 5.396-10.624 10.943-11.153v-5.972c-5.547-.529-8.934-4.649-10.943-11.153-2.002-6.484-2.28-14.437-2.066-20.577C105.214 5.894 100.233 0 93.5 0H24.508zM80 57.863C80 66.663 73.436 72 62.543 72H44a2 2 0 01-2-2V24a2 2 0 012-2h18.437c9.083 0 15.044 4.92 15.044 12.474 0 5.302-4.01 10.049-9.119 10.88v.277C75.317 46.394 80 51.21 80 57.863zM60.521 28.34H49.948v14.934h8.905c6.884 0 10.68-2.772 10.68-7.727 0-4.643-3.264-7.207-9.012-7.207zM49.948 49.2v16.458H60.91c7.167 0 10.964-2.876 10.964-8.281 0-5.406-3.903-8.178-11.425-8.178H49.948z"
                      fill="currentColor"></path>
            </svg>
            <span class="fs-4">Creeper 简易数据分析</span>
        </a>
    </header>

    <div class="row g-3 align-items-center">
        <div class="col-auto">
            <input type="text" class="form-control" id="key_world" placeholder="Key world">
        </div>
        <div class="col-auto">
            <label class="col-form-label">Index</label>
        </div>
        <div class="col-auto">
            <select class="form-select" aria-label="" id="index">
                <option selected>选择一个Index</option>
                {{$.Index}}
            </select>
        </div>

        <div class="col-auto">
            <input type="text" class="form-control" id="st" placeholder="开始时间  格式 20210101 可选">
        </div>

        <div class="col-auto">
            <input type="text" class="form-control" id="et" placeholder="结束时间  格式 20210101 可选">
        </div>

        <div class="col-auto">
            <input type="text" class="form-control" id="offset" placeholder="offset 可选">
        </div>

        <div class="col-auto">
            <input type="text" class="form-control" id="limit" placeholder="limit 可选">
        </div>

        <div class="col-auto">
            <button type="button" onclick="search()" class="btn btn-primary mb-3">查询</button>
        </div>
    </div>

    <table class="table table-striped table-hover">
        <thead>
        <tr>
            <th scope="col">CreateAt</th>
            <th scope="col">Message</th>
        </tr>
        </thead>
        <tbody id="tt">
			
        </tbody>
    </table>
</div>


<script>
    function search() {
        let keyWorld = $("#key_world").val()
        let startTime = $("#st").val()
        let endTime = $("#et").val()
        let index = $("#index option:selected").text()
        let limit = Number($("#limit").val())
        let offset = Number($("#offset").val())
        let token = $("#token").val()
        if (isNaN(limit) || index == "选择一个Index" || isNaN(offset)) {
            alert("参数错误  index 未选择 或 limit 非法")
            console.log(isNaN(limit))
            console.log(index)
            return
        }

        var requestURL = "/api/v1/web_search";
        var dataJSON = {};
        dataJSON["index"] = index;
        dataJSON["key_word"] = keyWorld;
        dataJSON["limit"] = limit;
        dataJSON["offset"] = offset;
        dataJSON["start_time"] = startTime;
        dataJSON["end_time"] = endTime;

        $.ajax({
            url: requestURL,
            data: JSON.stringify(dataJSON),
            type: "POST",
            dataType: "json",
            beforeSend: function (request) {
                request.setRequestHeader("token", token);
            },
            contentType: "application/json;charset=utf-8",
            success: function (returnData) {
                console.log(returnData);

                let r = ""
                returnData.data.data.forEach(function (rc) {
                    r += "<tr> <td>" + rc.create_at_string + "</td> <td>" + rc.message + "</td>"
                })
                $("#tt").html(r)
            },
            error: function (xhr, ajaxOptions, thrownError) {
                alert(xhr.responseText)
                console.log(xhr.status);
                console.log(thrownError);
            }
        });

        // console.log("k: ", keyWorld, " s: ", startTime, " e: ", endTime, " idx:", index, "  offset: ", offset)


    }
</script>
</body>
</html>
`

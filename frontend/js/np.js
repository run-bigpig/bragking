var $ = layui.jquery;
var baseUrl = "http://127.0.0.1:8080/";
layui.use(function () {
    getBaseData()
    getChartData()
    getTableData()
});

function getBaseData(){
    send(baseUrl + "api/findcloverbet", "post", {}, function (res) {
        $("#cloverbet").text(res.data)
    })
    send(baseUrl + "api/findmaxbetuser", "post", {user:"challenger",status:2}, function (res) {
        $("#maxbetchallenger").text(res.data.user+":"+res.data.count)
    })
    send(baseUrl + "api/findmaxbetuser", "post", {user:"contestant",status:1}, function (res) {
        $("#maxbetcontestant").text(res.data.user+":"+res.data.count)
    })
    send(baseUrl + "api/findmaxbetuser", "post", {user:"challenger",status:1}, function (res) {
        $("#minbetchallenger").text(res.data.user+":"+res.data.count)
    })
    send(baseUrl + "api/findmaxbetuser", "post", {user:"contestant",status:2}, function (res) {
        $("#minbetcontestant").text(res.data.user+":"+res.data.count)
    })
    send(baseUrl + "api/findmaxuser", "post", {user:"challenger",status:2}, function (res) {
        $("#maxchallenger").text(res.data.user+":"+res.data.count)
    })
    send(baseUrl + "api/findmaxuser", "post", {user:"contestant",status:1}, function (res) {
        $("#maxcontestant").text(res.data.user+":"+res.data.count)
    })
    send(baseUrl + "api/findmaxuser", "post", {user:"challenger",status:1}, function (res) {
        $("#maxfailchallenger").text(res.data.user+":"+res.data.count)
    })
    send(baseUrl + "api/findmaxuser", "post", {user:"contestant",status:2}, function (res) {
        $("#maxfailcontestant").text(res.data.user+":"+res.data.count)
    })
    send(baseUrl + "api/findcorrectanswercountlist", "post", {answer:1}, function (res) {
        $("#set1").text(res.data.count)
    })
    send(baseUrl + "api/findcorrectanswercountlist", "post", {answer:2}, function (res) {
        $("#set2").text(res.data.count)
    })
    send(baseUrl + "api/findmyanswercountlist", "post", {answer:1}, function (res) {
        $("#choose1").text(res.data.count)
    })
    send(baseUrl + "api/findmyanswercountlist", "post", {answer:2}, function (res) {
        $("#choose2").text(res.data.count)
    })
}

function chartInit(){
    let chartMap = new Map();
    let chartArr = ["setchart","choosechart"]
    for (let i = 0; i < 2; i++) {
        chartMap[chartArr[i]] = echarts.init(document.getElementById(chartArr[i]),myEchartsTheme);
        let option = {
            dataset:{
                source:[]
            },
            legend:{
                data: ['答案1','答案2']
            },
            tooltip: {
                trigger: 'axis',
                textStyle:{color:"#fff"}
            },
            grid: {
                left: '3%',
                right: '4%',
                bottom: '3%',
                containLabel: true
            },
            toolbox: {
                feature: {
                    saveAsImage: {}
                }
            },
            xAxis: {
                type: 'category',
                boundaryGap: false,
            },
            yAxis: {
                type: 'value'
            },
            series: [
                {name: '答案1', type: 'line'},
                {name: '答案2', type: 'line'},
            ]
        };
        chartMap[chartArr[i]].setOption(option);
    }
    return chartMap
}

function getChartData(){
    let chartMap = chartInit();
    console.log(chartMap)
    send(baseUrl + "api/finddatelist", "post", {name:"correct_answer"}, function (res) {
        chartMap["setchart"].setOption({
            dataset:{
                source:res.data
            }
        })
    })
    send(baseUrl + "api/finddatelist", "post", {name:"my_answer"}, function (res) {
        chartMap["choosechart"].setOption({
            dataset:{
                source:res.data
            }
        })
    })
    window.onresize = function () {
        chartMap["setchart"].resize();
        chartMap["choosechart"].resize();
    };
}

function tableInit(){
    var table = layui.table;
    let tablemap = new Map();
    let tablearr = ["tableset1","tableset2","tablechoose1","tablechoose2"]
    for (let i = 0; i < 4; i++) {
        tablemap.set(tablearr[i],table.render({
            elem: '#'+tablearr[i]
            ,data: []
            ,cols: [[
                {field:'user', title:'用户名'}
                ,{field:'win', title:'胜利次数'}
                ,{field:'fail', title:'失败次数'}
                ,{field:'total', title:'总次数'}
                ,{field:'winbet', title:'获胜妖晶'}
                ,{field:'failbet', title:'失败妖晶'}
                ,{field:'diffbet', title:'净胜妖精'}
                ,{field:'winrate', title:'胜率'}
            ]]
            ,page: true
        }));
    }
    return tablemap;
}

function getTableData(){
    let tablemap = tableInit();
    let fn = function (url,winwhere,failwhere,tablename) {
        send(url, "post", winwhere, function (winres) {
            send(url, "post", failwhere, function (failres) {
                let winarr = winres.data;
                let failarr = failres.data;
                let data = [];
                for (let i = 0; i < winarr.length; i++) {
                    let winuser = winarr[i].name;
                    let winwin = winarr[i].count;
                    let winbet = parseInt(winarr[i].bet*0.9);
                    let winfail = 0;
                    let failbet = 0;
                    for (let j = 0; j < failarr.length; j++) {
                        if (failarr[j].name === winuser) {
                            winfail = failarr[j].count;
                            failbet = failarr[j].bet;
                            break;
                        }
                    }
                    let wintotal = winwin + winfail;
                    let winrate = 0;
                    if (wintotal !== 0) {
                        winrate = (winwin / wintotal).toFixed(2);
                    }
                    data.push({
                        user: winuser,
                        win: winwin,
                        fail: winfail,
                        total: wintotal,
                        winrate: winrate
                        ,winbet: winbet,
                        failbet: failbet
                        ,diffbet: winbet - failbet
                    });
                }
                tablemap.get(tablename).reload({
                    data: data
                });
            })
        })
    }
    fn(baseUrl + "api/findchallengeranswercountlist", {answer:1,status:2}, {answer:1,status:1}, "tableset1");
    fn(baseUrl + "api/findchallengeranswercountlist", {answer:2,status:2}, {answer:2,status:1}, "tableset2");
    fn(baseUrl + "api/findcontestantmyanswercountlist", {answer:1,status:1}, {answer:1,status:2}, "tablechoose1");
    fn(baseUrl + "api/findcontestantmyanswercountlist", {answer:2,status:1}, {answer:2,status:2}, "tablechoose2");
}

function send(url, method, data, successCallback, errorCallback) {
    $.ajax({
        url: url,
        type: method,
        data: data,
        success: successCallback,
        error: errorCallback
    });
}
$().ready(() => {

    let appInfo = {}
    getAppInfo = (() => {
        fetch(ip + "/api/basic_setting/?app_id=gocr", {
            method: "GET",
        }).then((req) => {
            req.json(() => {
            }).then((msg) => {
                if (msg.code !== 200) {
                    // 0感叹号 1打钩 2打叉 3问号 4灰色的锁 5红色哭脸 6绿色笑脸
                    layer.msg(msg.message, {icon: 5})
                } else {
                    appInfo = msg.data
                    console.log(appInfo.app_name, "版本V" + appInfo.version)
                }
            }, () => {
                layer.msg('请求失败', {icon: 5});
            })
        }, () => {
            layer.msg('请求失败', {icon: 5});
        })
    })

    //init var
    g_result = ''
    g_result_join = ''
    g_image_type = 0

    //init func
    getAppInfo()

    layui.use('layer', function () {
        let layer = layui.layer;

        nav1 = (() => {
            // $("#content2").hide()
            // $("#content").show()
            $("legend").text("印刷文字识别")
            g_image_type = 0
        })

        nav2 = (() => {
            // $("#content").hide()K
            // $("#content2").show()
            $("legend").text("手写字体识别")
            g_image_type = 9
            // layer.open({
            //     title: '手写文字识别'
            //     , content: '正在开发中...'
            // });
        })

        // init
        nav1()

        nav_note = (() => {
            layer.open({
                title: '操作说明'
                ,
                content: '<div><li>1.单图识别：点击上传图片，然后点击开始识别按钮</li><li>2.多图识别：点击上传图片，选择多张图片，然后点击开始识别按钮</li><li>3.网络图片链接识别：复制图片地址，然后点击开始识别按钮</li><div>'
            });
        })


        nav_dev = (() => {
            layer.open({
                title: '开发日志'
                //, content: '<div>版本V0.1</div><div>功能：印刷体文字：单图识别，网络图片地址识别</div>'
                , content: appInfo['dev_log']
            });
        })

        nav_about = (() => {
            layer.open({
                title: '关于'
                , content: '<div>版本V0.1</div><div>开发者：fanbi </div><div>邮箱：unionline@126.com</div>'
                , content: appInfo['about_me']
            });

        })


        $('#result_copy').click(function () {
            let text = $("#result").val()
            $(this).attr('data-clipboard-text', text)
            var clipboard = new Clipboard(this);
            clipboard.on('success', function (e) {
                layer.msg('复制成功');
            });
            clipboard.on('error', function (e) {
                layer.msg('复制失败');
            });
        });


        //去除字符串中间空格
        function CTim(str) {
            return str.replace(/\s/g, '');
        }

        //去除换行
        function ClearBr(key) {
            key = key.replace(/<\/?.+?>/g, "");
            key = key.replace(/[\r\n]/g, "");
            return key;
        }

        $('#result_copy_join').click(function () {
            let text = $("#result").val()

            text = ClearBr(text)
            text = CTim(text)

            $(this).attr('data-clipboard-text', text)
            var clipboard = new Clipboard(this);
            clipboard.on('success', function (e) {
                layer.msg('复制成功');
            });
            clipboard.on('error', function (e) {
                layer.msg('复制失败');
            });
        });


        $('#result_join').click(function () {


            let flag = $('#result_join').attr("flag")

            if (flag == 0) {

                if (g_result_join == '') {

                    let text = ''
                    if (g_result == '') {
                        text = $("#result").val()
                        g_result = text
                    }
                    text = ClearBr(g_result)
                    // text = CTim(text)
                    g_result_join = text
                }else{

                }

                $('#result_join').attr("flag",1)
                $('#result_join').val("恢复换行")
                $("#result").val(g_result_join)

            } else {
                $('#result_join').attr("flag",0)
                $('#result_join').val("去除换行")
                $("#result").val(g_result)
            }

        });


        // 方法
        $("#result_clear").click(() => {
            $("#result").val("")
        })

    })


})


//注意：导航 依赖 element 模块，否则无法进行功能性操作
layui.use('element', function () {
    var element = layui.element;

    element.on('tab(tab)', function (data) {
        console.log(data);
    });

    element.on('nav(nav)', function (elem) {
        console.log(elem);
    });
    //…
});






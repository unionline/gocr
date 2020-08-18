$().ready(() => {

    layui.use(['upload', 'element', 'layer'], function () {
        // $ = layui.jquery;
        let upload = layui.upload;
        let layer = layui.layer;
        var element = layui.element;
        //拖拽上传
        var uploadInst = upload.render({
            elem: '#drag_upload'
            , url: prefix + '/upload/image/'
            , auto: false
            , bindAction: '#submit'
            , accept: 'images'
            , size: 1024*5 //KB
            , choose: function (obj) {
                obj.preview(function (index, file, result) {
                    $('#preview').attr('src', result); //图片链接（base64）
                    $('#preview').attr('min-width', '258px');
                    $('#preview').attr('height', '135px');
                    $('#img_box').attr('style', 'width: 258px;height: 135px;')
                    // $('#preview').attr('style', 'margin-top:2px;max-width:300px;min-width:180px;max-height:220px;');
                    $('#upload_icon').attr('style', 'visibility: hidden;font-size: 0px;');
                    // $('#drag_upload').attr('style', 'padding: 0px 0 0 0;height:220px')
                    $('#message').html('')
                });
            }
            , before: function (obj) { //obj参数包含的信息，跟 choose回调完全一致，可参见上文。
                layer.load(); //上传loading
            }, done: function (res, index, upload) {
                layer.closeAll('loading'); //关闭loading
                //如果上传失败
                if (res.code !== 200) {
                    layer.msg(res.message, {icon: 5})
                } else {
                    layer.msg('识别成功！', {icon: 1});
                    $("#result").val(res.data)
                }
            }
            , error: function () {
                layer.closeAll('loading'); //关闭loading
                //演示失败状态，并实现重传
                var demoText = $('#demoText');
                demoText.html('<a class="layui-btn layui-btn-normal demo-reload" style="float:left;margin-right: 15px"id="submit">重试</a>');
                demoText.find('.demo-reload').on('click', function () {
                    clear_result()
                    uploadInst.upload();
                });
            }
        });

        //多张图上传
        var uploadInst2 = upload.render({
            elem: '#upload_images'
            , url: prefix + '/upload/images/'
            , auto: false
            , bindAction: '#start_ocr_imgs'
            , accept: 'images'
            , multiple: true
            , size: 1024 //KB
            , number: 10
            , choose: function (obj) {

                obj.preview(function (index, file, result) {
                    let img = `<img src="` + result + `" alt="` + file.name + `" class="layui-upload-img" style="height:100px;padding-right: 10px">`
                    // img_list = img_list + img
                    $(".layui-upload-list").append(img)
                });
            }
            , progress: function (n, elem) {
                var percent = n + '%' //获取进度百分比
                element.progress('demo', percent); //可配合 layui 进度条元素使用

                //以下系 layui 2.5.6 新增
                console.log(elem); //得到当前触发的元素 DOM 对象。可通过该元素定义的属性值匹配到对应的进度条。
            }
            , allDone: function (obj) {

                console.log(obj.total); //得到总文件数
                console.log(obj.successful); //请求成功的文件数
                console.log(obj.aborted); //请求失败的文件数

            },
            done: function (res, index, upload) {
                layer.closeAll('loading'); //关闭loading
                //如果上传失败

                //如果上传失败
                if (res.code !== 200) {
                    layer.msg(res.message, {icon: 5})
                } else {
                    layer.msg('识别成功！', {icon: 1});

                    let ret = res.data
                    let cur_ret = $("#result").val()
                    $("#result").val(cur_ret + ret + "\n")

                }
            }
            , error: function () {
                //演示失败状态，并实现重传
                var demoText = $('#start_ocr_imgs');
                demoText.html('<a class="layui-btn layui-btn-normal demo-reload" style="float:left;margin-right: 15px"id="start_ocr_imgs">重试</a>');
                demoText.on('click', function () {
                    clear_result()
                    uploadInst2.upload();
                });
            }
        })
    })

    var clear_result = () => {
        $("#result").val("")
    }


    $("#start_ocr_img_url").click(() => {
        img_url = $("#img_url").val()
        img_url = encodeURIComponent(img_url)
        reqUrl = "/gocr/api/st/?img_url=" + img_url
        commonReq(reqUrl)
    })


    $("#reset_ocr_img_url").click(() => {
        $("#img_url").val("")
    })

    $("#reset_upload_image").click(() => {
        html = `  
                    <i id="upload_icon" class="layui-icon"></i>
                    <div id="img_box"><img class="layui-upload-img" id="preview"></div>
                    <p id="message" class="drag_message">点击上传，或将文件拖拽到此处</p>
                `
        setTimeout(() => {
            $('#drag_upload').html(html)
        }, 500)
    })

    $("#reset_upload_images").click(() => {
        html = `  
                <div class="layui-upload">
                    <button type="button" class="layui-btn" id="upload_images">多图片上传</button><input class="layui-upload-file" type="file" accept="undefined" name="file" multiple="">
                    <blockquote class="layui-elem-quote layui-quote-nm" style="margin-top: 10px;">
                        预览图：
                        <div class="layui-upload-list" id="preview_images"></div>
                    </blockquote>
                </div>`
        setTimeout(() => {
            $('#drag_upload2.layui-upload').html(html)
        }, 500)
    })


    var commonReq = (reqUrl, method, data) => {

        method == method || "get"

        fetch(reqUrl, {
            method: method,
            body: data
        }).then((req) => {
            req.json(() => {
            }).then((msg) => {
                if (msg.code !== 200) {
                    // 0感叹号 1打钩 2打叉 3问号 4灰色的锁 5红色哭脸 6绿色笑脸
                    layer.msg(msg.message, {icon: 5})
                } else {
                    $("#result").text(msg.data)
                    layer.msg("识别成功!", {icon: 1})
                }
            }, () => {
                layer.msg('请求失败', {icon: 5});
            })
        }, () => {
            layer.msg('请求失败', {icon: 5});
        })
    }

})





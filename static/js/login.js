/**
 * Created by fenggese on 2017/7/18.
 */
new Vue({
    el: '#main',
    data: {
        checked: true,
        ruleForm: {
            account: '',
            password: ''
        },
        rules: {
            account: [
                {required: true, message: '请输入账号', trigger: 'blur'},
            ],
            password: [
                {required: true, message: '请输入密码', trigger: 'blur'},
            ]
        },

    },
    methods: {
        submitForm() {
            axios({
                method: "POST",
                url: '/oa/v1/api/user/authenticate',
                data: {
                    account: this.ruleForm.account,
                    password: this.ruleForm.password
                }
            }).then(
                (rest) => {
                    if (rest.data.code === 200) {
                        // 登录成功 跳转到服务端传来的地址
                        window.location.href = "/";
                    } else {
                        // 登录失败
                        this.$message({
                            message: rest.data.msg,
                            type: 'error'
                        })
                    }
                });
        }
    },
})
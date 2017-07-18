/**
 * Created by fenggese on 2017/7/18.
 */
new Vue({
    el: "#main",
    data: {

    },
    methods: {
        logout() {
            window.location.href = "/oa/login";
            axios('/oa/v1/api/user/logout')
        }

    },
    computed: {
        defaultActive: function(){
            return '/';
        }
    },
});

var module = {}
var Admin = window.admin
var NewPage = window.newPage
var AdminMain = window.adminMain
var NewPage = window.newPage
var PageForm = window.pageForm
var ListPages = window.listPages

var appRoutes = {
    '/pages': function() {
        React.render(window.listPages(), document.querySelector("[data-ui='admin']"))
    }
};

var router = Router(appRoutes);

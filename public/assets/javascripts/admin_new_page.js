(function(){var l,d,a,b,f,i,h,k,g,j,e,c;g=typeof exports!=="undefined"&&exports!==null?exports:this;c=React.DOM,f=c.h1,d=c.button,h=c.label,e=c.ul,b=c.form,k=c.li,j=c.textarea,a=c.div,i=c.input;l=React.createClass({displayName:"NewPage",createPage:function(m){m.preventDefault();console.log("submitting");console.log(this.refs.title.getDOMNode().value);return $.ajax({type:"post",url:"/admin/pages/create",dataType:"json",data:{title:this.refs.title.getDOMNode().value,instructions:this.refs.instructions.getDOMNode().value,youtubes:this.refs.youtubes.getDOMNode().value,articles:this.refs.articles.getDOMNode().value,tags:this.refs.tags.getDOMNode().value},success:function(n){},done:function(n){return console.log(n)},fail:function(n){return console.log(n)}})},render:function(){return b({onSubmit:this.createPage,ref:"newPageForm"},f({},"Add New Page"),a({className:"row"},a({className:"form-group col-md-12"},h({},"Title"),i({className:"form-control",placeholder:"Title",ref:"title"})),a({className:"form-group col-md-6"},h({},"Instructions"),j({className:"form-control",placeholder:"Instructions, one per line",ref:"instructions"})),a({className:"form-group col-md-6"},h({},"Youtube Videos"),j({className:"form-control",placeholder:"Youtube links, one per line",ref:"youtubes"})),a({className:"form-group col-md-6"},h({},"Articles"),j({className:"form-control",placeholder:"Articles, one per line",ref:"articles"})),a({className:"form-group col-md-6"},h({},"Tags"),j({className:"form-control",placeholder:"Tags, one per line",ref:"tags"})),a({className:"form-group col-md-12"},d({className:"btn btn-primary col-md-12"},"Add Page"))))}});g.newPage=l}).call(this);
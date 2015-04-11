var Project = Backbone.Model.extend({
  urlRoot: 'http://localhost:1337/project',
  defaults: {
    image: 'images/new-project.jpg',
    name:'',
    description:'',
    link:'',
    author:'',
    tags:''
  },
  initialize: function(){}
});
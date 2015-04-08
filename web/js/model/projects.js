var Projects = Backbone.Collection.extend({
  model: Project,
  url: 'http://localhost:1337/projects'
});
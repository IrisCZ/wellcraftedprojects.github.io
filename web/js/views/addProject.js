AddProjectView = Backbone.View.extend({

  model: new Project(),
  template: _.template($('#new_project_template').html()),

  events: {
    'click #newproject-save':'submit',
    'click #new-logo':'openFileSelector',
    'change #new-logo-input':'showPreview'
  },

  initialize: function(){
  },

  render: function(){
    this.$el.append(this.template({}));
    this.modalWindow = this.$('#newproject')
    return this;
  },

  submit: function(){
    this.model.set('name', this.$('#new-name').val().trim());
    this.model.set('url', this.$('#new-url').val().trim());
    this.model.set('author', this.$('#new-author').val().trim());
    this.model.set('tags', this.$('#new-tags').val().trim());
    this.model.set('description', this.$('#new-description').val().trim());
    this.model.set('image', this.$('#new-logo .imagePreview img').attr('src'));
    var self = this;
    this.model.save(null, {
      success: function(model, response){
        model.clear().set(model.defaults);
        self.modalWindow.modal('hide');
        Backbone.trigger('project:created');
      }
    });
  },

  openFileSelector: function(){
    this.$('#new-logo-input').trigger("click");
  },

  showPreview: function(){
    previewImage(this.$('#new-logo-input')[0],[1],4);
  }
});
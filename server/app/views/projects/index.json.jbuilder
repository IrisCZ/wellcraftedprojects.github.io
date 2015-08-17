json.array!(@projects) do |project|
  json.extract! project, :id, :name, :author, :url, :description, :positives, :negatives, :image
  json.url project_url(project, format: :json)
end

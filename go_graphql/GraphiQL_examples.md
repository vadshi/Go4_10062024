```
# Welcome to GraphiQL
#
# GraphiQL is an in-browser tool for writing, validating, and
# testing GraphQL queries.
#
# Type queries into this side of the screen, and you will see intelligent
# typeaheads aware of the current GraphQL type schema and live syntax and
# validation errors highlighted within the text.
#
# GraphQL queries typically start with a "{" character. Lines that start
# with a # are ignored.
#
# An example GraphQL query might look like:
#
#     {
#       field(arg: "value") {
#         subField
#       }
#     }
#
# Keyboard shortcuts:
#
#   Prettify query:  Shift-Ctrl-P (or press the prettify button)
#
#  Merge fragments:  Shift-Ctrl-M (or press the merge button)
#
#        Run Query:  Ctrl-Enter (or press the play button)
#
#    Auto Complete:  Ctrl-Space (or just start typing)
#
mutation CreatePost {
  CreatePost(
    input: {
      Title: "How to create new GraphQL app", 
      Content: "Мы создаем новый вид API с помощью GraphQL",
      Author: "User",
      Hero: "User picture link"
    }
  )
  {
    _id
    Title
    Author
  }
}

query GetOnePost {
  GetOnePost(id: "6666ba378d120a3e0006d8df") {
    _id
    Title
    Content
    Author
    Hero
    Published_At
    Updated_At
  }
}

query GetAllPosts {
  GetAllPosts {
    _id
    Title
    Content
    Author
    Hero
    Published_At
    Updated_At
  }
}

mutation UpdatePost{
  UpdatePost(id: "6666ba378d120a3e0006d8df"
  input: {
    Title: "Change title to check of working",
    Content: "Change content"
  }
  )
  {
    _id
    Title
    Content
  }
}

mutation DeletePost {
  DeletePost(id: "6666ba378d120a3e0006d8df") {
    deletePostId
  }
}


query Types {
__type(name: "Post") {
    fields {
      name
    }
  }
}
```
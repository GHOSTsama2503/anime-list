package anilist

const QuerySearch = `
query ($id: Int, $page: Int, $perPage: Int, $search: String) {
  Page (page: $page, perPage: $perPage) {
    pageInfo {
      total
      currentPage
      lastPage
      hasNextPage
      perPage
    }
    media (id: $id, search: $search) {
      id
      title {
        romaji
      }
      coverImage {
        extraLarge
        large
        medium
        color
      }
      bannerImage
    }
  }
}
`

const QueryInfo = `
query ($id: Int) {
  Media (id: $id, type: ANIME) {
    id
    title {
      romaji
    }
    coverImage {
      extraLarge
      large
      medium
      color
    }
    bannerImage
    description
  }
}
`

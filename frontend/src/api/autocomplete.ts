export type Suggestion = {
  value: string
  score: number
}

type GraphQLError = {
  message: string
}

type AutocompleteResponse = {
  data?: {
    autocomplete: Suggestion[]
  }
  errors?: GraphQLError[]
}

const autocompleteQuery = `
  query Autocomplete($prefix: String!) {
    autocomplete(prefix: $prefix) {
      value
      score
    }
  }
`

export async function searchSuggestions(
  prefix: string,
  signal?: AbortSignal,
): Promise<Suggestion[]> {
  const response = await fetch('/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      query: autocompleteQuery,
      variables: {
        prefix,
      },
    }),
    signal,
  })

  if (!response.ok) {
    throw new Error(`HTTP request failed with status ${response.status}`)
  }

  const payload = (await response.json()) as AutocompleteResponse

  if (payload.errors && payload.errors.length > 0) {
    throw new Error(payload.errors[0].message)
  }

  if (!payload.data) {
    throw new Error('GraphQL response did not include data')
  }

  return payload.data.autocomplete
}
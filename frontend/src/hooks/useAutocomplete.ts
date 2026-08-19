import { useEffect, useState } from 'react'
import {
  searchSuggestions,
  type Suggestion,
} from '../api/autocomplete'

const SEARCH_DEBOUNCE_MS = 300

export type UseAutocompleteResult = {
  prefix: string
  setPrefix: (value: string) => void
  selectPrefix: (value: string) => void
  isSearchEnabled: boolean
  normalizedPrefix: string
  suggestions: Suggestion[]
  isLoading: boolean
  error: string | null
}

export function useAutocomplete(
  debounceMs: number = SEARCH_DEBOUNCE_MS,
): UseAutocompleteResult {
  const [prefix, setPrefixState] = useState('')
  const [shouldSearch, setShouldSearch] = useState(true)
  const [suggestions, setSuggestions] = useState<Suggestion[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const normalizedPrefix = prefix.trim().toLowerCase()

  const setPrefix = (value: string) => {
    setShouldSearch(true)
    setPrefixState(value)
  }

  const selectPrefix = (value: string) => {
    setShouldSearch(false)
    setPrefixState(value)
  }

  useEffect(() => {
    if (!shouldSearch) {
      setSuggestions([])
      setIsLoading(false)
      setError(null)
      return
    }

    if (normalizedPrefix === '') {
      setSuggestions([])
      setIsLoading(false)
      setError(null)
      return
    }

    const controller = new AbortController()

    setSuggestions([])
    setIsLoading(true)
    setError(null)

    const timeoutId = window.setTimeout(() => {
      searchSuggestions(normalizedPrefix, controller.signal)
        .then((results) => {
          setSuggestions(results)
        })
        .catch((requestError: unknown) => {
          if (
            requestError instanceof DOMException &&
            requestError.name === 'AbortError'
          ) {
            return
          }

          setSuggestions([])

          setError(
            requestError instanceof Error
              ? requestError.message
              : 'An unexpected error occurred',
          )
        })
        .finally(() => {
          if (!controller.signal.aborted) {
            setIsLoading(false)
          }
        })
    }, debounceMs)

    return () => {
      window.clearTimeout(timeoutId)
      controller.abort()
    }
  }, [normalizedPrefix, debounceMs, shouldSearch])

  return {
    prefix,
    setPrefix,
    selectPrefix,
    isSearchEnabled: shouldSearch,
    normalizedPrefix,
    suggestions,
    isLoading,
    error,
  }
}

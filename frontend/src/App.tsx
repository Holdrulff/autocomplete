import { useState } from 'react'
import './App.css'

type Suggestion = {
  value: string
  score: number
}

const previewSuggestions: Suggestion[] = [
  { value: 'javascript', score: 2_522_113 },
  { value: 'java', score: 1_914_698 },
  { value: 'javafx', score: 38_812 },
  { value: 'reactjs', score: 473_924 },
  { value: 'python', score: 2_205_065 },
]

const scoreFormatter = new Intl.NumberFormat('en-US')

function App() {
  const [prefix, setPrefix] = useState('')
  const normalizedPrefix = prefix.trim().toLowerCase()

  const suggestions =
    normalizedPrefix === ''
      ? []
      : previewSuggestions.filter((suggestion) =>
          suggestion.value.startsWith(normalizedPrefix),
        )

  return (
    <main className="app-shell">
      <section
        className="search-card"
        aria-labelledby="search-title"
      >
        <p className="eyebrow">Autocomplete</p>

        <h1 id="search-title">
          Find a technology
        </h1>

        <p className="search-description">
          Start typing to explore popular technology tags from Stack Overflow.
        </p>

        <div className="search-field">
          <label htmlFor="autocomplete-input">
            Technology
          </label>

          <input
            id="autocomplete-input"
            name="autocomplete"
            type="search"
            placeholder="e.g., java, react, python"
            autoComplete="off"
            value={prefix}
            onChange={(event) => setPrefix(event.target.value)}
            aria-describedby="search-help"
          />
        </div>
        <p
          id="search-help"
          className="search-hint"
          aria-live="polite"
        >
          {normalizedPrefix === ''
            ? 'Type at least one character.'
            : suggestions.length === 0
              ? 'No suggestions found.'
              : `${suggestions.length} ${
                suggestions.length === 1 ? 'suggestion' : 'suggestions'
              } found.`}
        </p>
        {suggestions.length > 0 && (
          <ul
            className="suggestion-list"
            aria-label="Autocomplete suggestions"
          >
            {suggestions.map((suggestion) => (
              <li key={suggestion.value}>
                <button
                  type="button"
                  className="suggestion-item"
                  onClick={() => setPrefix(suggestion.value)}
                >
                  <span className="suggestion-value">
                    {suggestion.value}
                  </span>

                  <span className="suggestion-score">
                    {scoreFormatter.format(suggestion.score)} questions
                  </span>
                </button>
              </li>
            ))}
          </ul>
        )}
      </section>
    </main>
  )
}

export default App

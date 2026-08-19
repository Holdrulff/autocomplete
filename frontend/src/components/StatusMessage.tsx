import { WarningIcon } from './icons'
import './StatusMessage.css'

type StatusMessageProps = {
  id: string
  normalizedPrefix: string
  isLoading: boolean
  error: string | null
  resultCount: number
}

export function StatusMessage({
  id,
  normalizedPrefix,
  isLoading,
  error,
  resultCount,
}: StatusMessageProps) {
  const hintText =
    normalizedPrefix === ''
      ? 'Type at least one character.'
      : isLoading
        ? 'Loading suggestions...'
        : error
          ? 'Unable to load suggestions.'
          : resultCount === 0
            ? 'No suggestions found.'
            : `${resultCount} ${resultCount === 1 ? 'suggestion' : 'suggestions'} found.`

  return (
    <>
      <p id={id} className="search-hint" aria-live="polite">
        {hintText}
      </p>

      {error && (
        <p className="search-error" role="alert">
          <WarningIcon />
          <span>{error}</span>
        </p>
      )}
    </>
  )
}

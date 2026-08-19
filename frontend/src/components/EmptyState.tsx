import { NoResultsIcon } from './icons'

type EmptyStateProps = {
  normalizedPrefix: string
}

export function EmptyState({ normalizedPrefix }: EmptyStateProps) {
  return (
    <div className="empty-state" aria-hidden="true">
      <NoResultsIcon />
      <span>No matches for &ldquo;{normalizedPrefix}&rdquo;. Try a different spelling.</span>
    </div>
  )
}

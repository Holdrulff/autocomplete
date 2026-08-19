import type { SVGProps } from 'react'

function BaseIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="20"
      height="20"
      viewBox="0 0 20 20"
      fill="none"
      stroke="currentColor"
      strokeWidth="1.6"
      strokeLinecap="round"
      strokeLinejoin="round"
      aria-hidden="true"
      focusable="false"
      {...props}
    />
  )
}

export function SearchIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <BaseIcon {...props}>
      <circle cx="8.5" cy="8.5" r="5.5" />
      <line x1="16.5" y1="16.5" x2="12.6" y2="12.6" />
    </BaseIcon>
  )
}

export function ClearIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <BaseIcon {...props}>
      <line x1="5" y1="5" x2="15" y2="15" />
      <line x1="15" y1="5" x2="5" y2="15" />
    </BaseIcon>
  )
}

export function SunIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <BaseIcon {...props}>
      <circle cx="10" cy="10" r="3.5" />
      <line x1="10" y1="1.5" x2="10" y2="3.5" />
      <line x1="10" y1="16.5" x2="10" y2="18.5" />
      <line x1="1.5" y1="10" x2="3.5" y2="10" />
      <line x1="16.5" y1="10" x2="18.5" y2="10" />
      <line x1="4.2" y1="4.2" x2="5.6" y2="5.6" />
      <line x1="14.4" y1="14.4" x2="15.8" y2="15.8" />
      <line x1="4.2" y1="15.8" x2="5.6" y2="14.4" />
      <line x1="14.4" y1="5.6" x2="15.8" y2="4.2" />
    </BaseIcon>
  )
}

export function MoonIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <BaseIcon {...props}>
      <path d="M17 11.2A7 7 0 1 1 8.8 3a5.6 5.6 0 0 0 8.2 8.2z" />
    </BaseIcon>
  )
}

export function WarningIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <BaseIcon {...props}>
      <path d="M10 3.2 17.5 16H2.5Z" />
      <line x1="10" y1="8" x2="10" y2="11.5" />
      <line x1="10" y1="13.8" x2="10" y2="13.9" />
    </BaseIcon>
  )
}

export function NoResultsIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <BaseIcon {...props} strokeDasharray="2.4 2.4">
      <circle cx="8.5" cy="8.5" r="5.5" />
      <line x1="16.5" y1="16.5" x2="12.6" y2="12.6" strokeDasharray="none" />
    </BaseIcon>
  )
}

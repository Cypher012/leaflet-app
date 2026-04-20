// src/hooks/useTheme.ts
import { useState, useEffect } from 'react'

type ThemeMode = 'light' | 'dark'

function getInitialMode(): ThemeMode {
  if (typeof window === 'undefined') return 'light'
  const stored = window.localStorage.getItem('theme')
  if (stored === 'light' || stored === 'dark') return stored
  return window.matchMedia('(prefers-color-scheme: dark)').matches
    ? 'dark'
    : 'light'
}

function applyThemeMode(mode: ThemeMode) {
  document.documentElement.classList.remove('light', 'dark')
  document.documentElement.classList.add(mode)
  document.documentElement.setAttribute('data-theme', mode)
  document.documentElement.style.colorScheme = mode
  window.localStorage.setItem('theme', mode)
}

export function useTheme() {
  const [mode, setMode] = useState<ThemeMode>('light')
  const [spinning, setSpinning] = useState(false)

  useEffect(() => {
    const initial = getInitialMode()
    setMode(initial)
    applyThemeMode(initial)
  }, [])

  function toggle() {
    setSpinning(true)
    setTimeout(() => setSpinning(false), 400)
    const next: ThemeMode = mode === 'light' ? 'dark' : 'light'
    setMode(next)
    applyThemeMode(next)
  }

  return { mode, spinning, toggle }
}

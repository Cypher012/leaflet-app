// src/components/web/ThemeToggle.tsx
import { Sun, Moon } from 'lucide-react'
import { Switch } from '../ui/switch'
import { useTheme } from '#/hooks/useTheme'

export default function ThemeToggle() {
  const { mode, spinning, toggle } = useTheme()
  const Icon = mode === 'dark' ? Moon : Sun

  return (
    <>
      <style>{`
        @keyframes spin-once {
          0%   { transform: rotate(0deg) scale(1); }
          40%  { transform: rotate(180deg) scale(0.75); }
          100% { transform: rotate(360deg) scale(1); }
        }
        .icon-spin { animation: spin-once 0.4s cubic-bezier(0.34, 1.56, 0.64, 1) forwards; }
      `}</style>
      <div
        onClick={toggle}
        className="md:text-sm text-xs text-foreground/80 hover:bg-muted hover:text-foreground flex w-full cursor-pointer items-center space-x-3 rounded-lg px-3 py-2 font-medium transition-colors"
      >
        <Icon className={`md:size-4 size-3.5 ${spinning ? 'icon-spin' : ''}`} />
        <span>{mode === 'dark' ? 'Dark' : 'Light'} Mode</span>
        <Switch
          checked={mode === 'dark'}
          className="ml-auto pointer-events-none data-[state=unchecked]:bg-muted-foreground/30"
        />
      </div>
    </>
  )
}

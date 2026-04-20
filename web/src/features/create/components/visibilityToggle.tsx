import { useState } from 'react'
import { motion } from 'framer-motion'
import { cn } from '#/lib/utils'

type Visibility = 'public' | 'private'

interface VisibilityToggleProps {
  value?: Visibility
  onChange?: (value: Visibility) => void
}

export default function VisibilityToggle({
  value,
  onChange,
}: VisibilityToggleProps) {
  const [selected, setSelected] = useState<Visibility>(value ?? 'public')

  function select(v: Visibility) {
    setSelected(v)
    onChange?.(v)
  }

  return (
    <div className="relative flex items-center rounded-full bg-background p-2.5">
      {(['public', 'private'] as Visibility[]).map((v) => (
        <button
          key={v}
          onClick={() => select(v)}
          className="relative z-10 px-4 py-2 text-xs font-semibold uppercase tracking-wide transition-colors"
        >
          {selected === v && (
            <motion.div
              layoutId="visibility-pill"
              className="absolute inset-0 rounded-full bg-primary"
              transition={{ type: 'spring', stiffness: 400, damping: 30 }}
            />
          )}
          <span
            className={cn(
              'relative z-10 transition-colors',
              selected === v
                ? 'text-primary-foreground'
                : 'text-muted-foreground',
            )}
          >
            {v}
          </span>
        </button>
      ))}
    </div>
  )
}

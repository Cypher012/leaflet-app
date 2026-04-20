'use client'

import { cn } from '#/lib/utils'
import { buttonVariants } from '@/components/ui/button'
import { Link } from '@tanstack/react-router'

export default function Hero() {
  return (
    <div className="bg-[#eff3f1] dark:bg-background w-full overflow-hidden">
      <div className="container mx-auto px-4 py-24 sm:px-6 lg:px-8 lg:py-32 lg:pb-40">
        <div className="mx-auto max-w-3xl ">
          {/* Heading */}
          <h1 className="text-primary bg-clip-text text-center font-semibold text-4xl tracking-tighter text-balance sm:text-5xl md:text-6xl lg:text-8xl">
            Thought worth Keeping.
          </h1>

          {/* Description */}
          <p className="text-muted-foreground mx-auto mt-6 max-w-xl text-center text-lg leading-7 text-balance">
            Write short form notes. Share them publicly. Discover ideas worth
            reading in a focused, botanical digital garden.
          </p>

          {/* CTA Buttons */}
          <div className="mt-10 flex flex-col items-center justify-center gap-4 sm:flex-row">
            <Link
              to="/"
              className={cn(
                buttonVariants({ variant: 'secondary', size: 'lg' }),
                'group text-lg hover:shadow-primary/30 relative overflow-hidden rounded-full py-7 px-9 shadow-lg transition-all duration-300',
              )}
            >
              Start writing
            </Link>

            <Link
              to="/feeds"
              className={cn(
                buttonVariants({ variant: 'secondary', size: 'lg' }),
                'flex items-center gap-2 rounded-full backdrop-blur-sm text-lg py-7 px-9',
              )}
            >
              Browse cards
            </Link>
          </div>
        </div>
      </div>
    </div>
  )
}

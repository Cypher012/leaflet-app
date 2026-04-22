import { cn } from '#/lib/utils'
import { Leaf } from 'lucide-react'


type LikeButtonProps = {
    isLiked: boolean,
    likes : number,
    handleLike:  () => Promise<void>
}

const LikeButton = ({handleLike, isLiked, likes}: LikeButtonProps) => {
  return (
      <button
    onClick={(e) => {
      e.preventDefault()
      e.stopPropagation()
      handleLike()
    }}
    className="flex items-center gap-2"
  >
    <Leaf className={cn('md:size-4 size-3.5', isLiked ? 'text-primary fill-primary' : '')} />
    <span className='md:text-sm text-xs'>{likes}</span>
  </button>
  )
}

export default LikeButton
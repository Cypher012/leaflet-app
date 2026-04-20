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
    <Leaf className={cn('w-4 h-4', isLiked ? 'text-primary fill-primary' : '')} />
    <span>{likes}</span>
  </button>
  )
}

export default LikeButton
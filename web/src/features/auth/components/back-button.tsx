import { Button } from '#/components/ui/button'
import { useNavigate } from '@tanstack/react-router'
import { ArrowLeft } from 'lucide-react'

const BackButton = () => {
  const navigate = useNavigate()
  return (
    <Button
      variant="ghost"
      onClick={() => navigate({ to: '..' })}
      className="group absolute top-6 left-6 flex items-center gap-2 px-3 py-2 rounded-lg text-primary"
    >
      <ArrowLeft className="w-4 h-4 transition-all duration-200 ease-out group-hover:-translate-x-1.5" />
      <span className="text-sm">Back</span>
    </Button>
  )
}

export default BackButton

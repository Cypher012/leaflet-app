// src/components/ui/FileUpload.tsx
import { useState, useRef } from 'react'
import { cn } from '#/lib/utils'
import { Upload, X } from 'lucide-react'
import { PlantBucket } from '#/icons/plantBucket.icon'

interface FileUploadProps {
  onFile: (file: File) => void
  accept?: string
  label?: string
}

export default function FileUpload({
  onFile,
  accept = 'image/*',
  label = 'Upload file',
}: FileUploadProps) {
  const [isDragging, setIsDragging] = useState(false)
  const [preview, setPreview] = useState<string | null>(null)
  const inputRef = useRef<HTMLInputElement>(null)

  function handleFile(file: File) {
    onFile(file)
    if (file.type.startsWith('image/')) {
      setPreview(URL.createObjectURL(file))
    }
  }

  function clearFile() {
    setPreview(null)
    if (inputRef.current) inputRef.current.value = ''
  }

  function onDrop(e: React.DragEvent) {
    e.preventDefault()
    setIsDragging(false)
    handleFile(e.dataTransfer.files[0])
  }

  if (preview) {
    return (
      <div className="relative aspect-video max-w-2xl mx-auto rounded-xl overflow-hidden border-2 border-green-700">
        <img
          src={preview}
          className="absolute inset-0 h-full w-full object-cover"
        />
        <button
          onClick={clearFile}
          className="absolute top-2 right-2 rounded-full bg-black/60 p-1.5 text-white hover:bg-black/80 transition-colors"
        >
          <X className="h-4 w-4" />
        </button>
      </div>
    )
  }

  return (
    <div
      onClick={() => inputRef.current?.click()}
      onDragOver={(e) => {
        e.preventDefault()
        setIsDragging(true)
      }}
      onDragLeave={() => setIsDragging(false)}
      onDrop={onDrop}
      className={cn(
        'relative flex flex-col items-center bg-background h-80 justify-center gap-3 rounded-xl border-2 border-dashed p-8 cursor-pointer transition-colors',
        isDragging
          ? 'border-primary bg-primary/5'
          : 'border-border/50 hover:border-primary/50 hover:bg-muted/50',
      )}
    >
      <input
        ref={inputRef}
        type="file"
        accept={accept}
        className="hidden"
        onChange={(e) => {
          const file = e.target.files?.[0]
          if (file) handleFile(file)
        }}
      />
      <div>
        <PlantBucket className="size-10 text-muted-foreground" />
      </div>
      <div className="text-center">
        <p className="text-sm text-muted-foreground font-medium">{label}</p>
      </div>
    </div>
  )
}

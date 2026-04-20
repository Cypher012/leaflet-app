import { Textarea } from '#/components/ui/textarea'
import { CircleAlert, Tag } from 'lucide-react'
import FileUpload from './fileUpload'
import { useForm } from '@tanstack/react-form'
import { Button } from '#/components/ui/button'
import z from 'zod'
import { cn } from '#/lib/utils'
import { useCreateFeed } from '../api/mutations'
import { uploadFile } from '../api/api'
import { useState } from 'react'
import { useRouter } from '@tanstack/react-router'

export const createCardSchema = z.object({
  title: z.string().min(1, 'Please fill in this field.').max(200),
  content: z.string(),
  feedImage: z.instanceof(File).nullable(),
})

const CreatePage = () => {
  const { mutateAsync, isPending } = useCreateFeed()
  const [fileKey, setFileKey] = useState(0)
  const [isUploading, setIsUploading] = useState(false)

  const router = useRouter()

  const form = useForm({
    defaultValues: {
      title: '',
      content: '',
      feedImage: null as File | null,
    },
    validators: {
      onSubmit: createCardSchema,
    },
    onSubmit: async ({ value }) => {
      let feedImageURL: string | null = null

      if (value.feedImage) {
        try {
          setIsUploading(true)
          feedImageURL = await uploadFile(value.feedImage, 'feed')
        } finally {
          setIsUploading(false)
        }
      }

      await mutateAsync({
        title: value.title,
        content: value.content,
        feed_image: feedImageURL,
      })

      form.reset()
      setFileKey((k) => k + 1)
    },
  })

  const isBusy = isPending || form.state.isSubmitting || isUploading

  return (
    <div className="section-container w-full max-w-3xl pb-20">
      <form
        onSubmit={(e) => {
          e.preventDefault()
          form.handleSubmit()
        }}
      >
        <div className="w-full min-h-200 rounded-4xl p-10 overflow-hidden bg-card">
          <div className="space-y-10">
            <form.Field name="feedImage">
              {(field) => (
                <FileUpload
                  key={fileKey}
                  onFile={(file) => field.handleChange(file)}
                  accept="image/*"
                  label="Capture the botanical details - Drag & drop or click to upload your find."
                />
              )}
            </form.Field>

            <div className="space-y-4">
              <div>
                <form.Field name="title">
                  {(field) => (
                    <>
                      <Textarea
                        value={field.state.value}
                        onChange={(e) =>
                          field.handleChange(e.target.value.slice(0, 150))
                        }
                        onBlur={field.handleBlur}
                        placeholder="Untitled Observation..."
                        className={cn(
                          'shadow-none font-bold p-3 text-3xl! min-h-12! rounded-sm! text-muted-foreground placeholder:text-primary/60 border resize-none overflow-hidden',
                          field.state.meta.errors.length
                            ? 'border-destructive/30'
                            : 'border-border/20',
                        )}
                      />
                      <div className="flex justify-between mt-2 text-[13px] font-medium">
                        <span className="text-destructive flex font-medium">
                          {field.state.meta.errors.length ? (
                            <CircleAlert className="mr-2 size-4" />
                          ) : (
                            ''
                          )}{' '}
                          {field.state.meta.errors[0]?.message}
                        </span>
                        <span
                          className={
                            field.state.value.length >= 150
                              ? 'text-destructive'
                              : 'text-muted-foreground'
                          }
                        >
                          {field.state.value.length}/150
                        </span>
                      </div>
                    </>
                  )}
                </form.Field>
              </div>
              <form.Field name="content">
                {(field) => (
                  <Textarea
                    value={field.state.value}
                    onChange={(e) => field.handleChange(e.target.value)}
                    onBlur={field.handleBlur}
                    placeholder="Describe the ephemeral details of your find..."
                    className="shadow-none text-base! p-4 rounded-sm! text-muted-foreground h-68 placeholder:text-primary/60 border border-border/20 resize-none"
                  />
                )}
              </form.Field>
            </div>

            <div className="flex justify-between items-center">
              <div className="bg-secondary text-primary py-2.5 px-4 font-semibold text-[13px] flex rounded-4xl">
                <Tag className="mr-3 size-4" /> ADD TAG
              </div>
            </div>
          </div>
        </div>
        <div className="mt-10 flex items-center">
          <Button
            onClick={() => router.history.back()}
            variant="ghost"
            className="mr-auto py-6 px-5 rounded-4xl"
          >
            Cancel
          </Button>
          <div className="flex items-center gap-2">
            <Button
              type="submit"
              variant="default"
              disabled={isPending || form.state.isSubmitting}
              className="py-6 px-5 rounded-4xl text-gray-900"
            >
              {isBusy ? (
                <span className="animate-pulse">Publishing...</span>
              ) : (
                'Publish Card'
              )}
            </Button>
          </div>
        </div>
      </form>
    </div>
  )
}

export default CreatePage

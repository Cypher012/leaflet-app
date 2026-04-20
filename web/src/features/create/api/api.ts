import { API_ROUTES } from '#/lib/api-routes'
import { apiClient } from '#/lib/axios'
import type { ApiResponse } from '#/types/global'

type UploadType = 'feed' | 'avatar' | 'comment'

type UploadPresignResponse = {
  public_url: string
  upload_url: string
}

export async function uploadFile(
  file: File,
  type: UploadType,
): Promise<string> {
  console.log(API_ROUTES.upload.presign)
  const { data } = await apiClient.get<ApiResponse<UploadPresignResponse>>(
    API_ROUTES.upload.presign,
    {
      params: {
        type,
        content_type: file.type,
      },
    },
  )

  const r2Response: Response = await fetch(data.upload_url, {
    method: 'PUT',
    body: file,
    headers: { 'Content-Type': file.type },
  })

  if (!r2Response.ok) {
    throw new Error(`Upload failed: ${r2Response.statusText}`)
  }

  return data.public_url
}

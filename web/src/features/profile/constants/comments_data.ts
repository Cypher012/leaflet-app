type MOCK_COMMENTS_TYPE = {
  id: string
  postTitle: string
  commentBody: string
  timestamp: string
  likes: number
}[]

export const MOCK_COMMENTS: MOCK_COMMENTS_TYPE = [
  {
    id: '1',
    postTitle: 'MIDNIGHT BLOOM STUDY',
    commentBody:
      'The way you captured the lunar luminescence on the petals is extraordinary. It reminds me of the 18th-century botanical sketches by Maria Sibylla Merian.',
    timestamp: '5h ago',
    likes: 12,
  },
  {
    id: '2',
    postTitle: 'FERN STRUCTURES VOL. 4',
    commentBody:
      'The symmetry in the fronds is mathematically perfect. I love how the dark background makes the spore patterns pop.',
    timestamp: '12h ago',
    likes: 8,
  },
  {
    id: '3',
    postTitle: 'DESERT XEROPHYTE',
    commentBody:
      'Brilliant use of texture. You can almost feel the waxy coating on the succulents. Keep up the amazing work!',
    timestamp: '2d ago',
    likes: 24,
  },
  {
    id: '4',
    postTitle: 'MOSS TAXONOMY',
    commentBody:
      'Finally, someone giving Bryophytes the spotlight they deserve! The macro detail here is simply unmatched.',
    timestamp: '1w ago',
    likes: 15,
  },
]

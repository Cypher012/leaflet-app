const useOAuth = () => {
  const loginWithGitHub = () => {
    // Use import.meta.env and the VITE_ prefix
    const url = import.meta.env.VITE_OAUTH_GITHUB_LOGIN
    if (url) {
      window.location.href = url
    } else {
      console.error('GitHub Login URL is undefined. Check your .env file!')
    }
  }

  const loginWithGoogle = () => {
    const url = import.meta.env.VITE_OAUTH_GOOGLE_LOGIN
    if (url) {
      window.location.href = url
    } else {
      console.error('Google Login URL is undefined. Check your .env file!')
    }
  }

  return { loginWithGitHub, loginWithGoogle }
}

export default useOAuth

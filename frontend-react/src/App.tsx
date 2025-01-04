import { useGoogleLogin } from '@react-oauth/google';
import axios from "axios";

function App() {

  const login = useGoogleLogin({
    onSuccess: async (resp) => {
      console.log(resp);

      const tokenResponse = await axios.get(
        `http://localhost:8080/auth/google/callback?code=${resp.code}`
      )

      console.log("JWT Token", tokenResponse);
    },
    flow: 'auth-code',
  })

  return (
    <div style={{width: '100vw', height: '100vh', display: 'flex', flexDirection: 'column', justifyContent: 'center', alignItems: 'center'}}>
      <h1>Hello World!</h1>
      <div>
        <button onClick={() => login()}>Sign In with Google</button>
      </div>
    </div>
  )
}

export default App

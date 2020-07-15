import axios from 'axios'
import { ThunkActionSync } from '~/types/actions'
import Firebase from '~/libs/Firebase'

type GoogleLoginPayloads = {
    uid: string
    email: string
    full_name: string
    username: string
    token: string
}

export const googleLogin = (isSignedIn: boolean): ThunkActionSync => async dispatch => {
    const { auth } = Firebase
    const provider = new auth.GoogleAuthProvider()
    provider.addScope('https://www.googleapis.com/auth/contacts.readonly')

    try {
        const { user } = await auth().signInWithPopup(provider)
        const token = await auth().currentUser?.getIdToken()

        if (token) {
            console.log(token)
            const res = await axios.post(
                'http://localhost:8088/accounts/signup',
                { uid: user?.uid },
                {
                    headers: { 'Authorization': `Bearer ${token}` }
                }
            )
            console.log(res)
        } else {
            // dispatch to some error state
        }
    } catch (err) {
        console.error(err)
        console.log(err)
    }
}

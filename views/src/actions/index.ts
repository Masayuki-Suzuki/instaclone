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
        const result = await auth().signInWithPopup(provider)
        const token = await auth().currentUser?.getIdToken()
        console.log(result)
        console.log(token)
    } catch (err) {
        console.error(err)
        console.log(err)
    }
}

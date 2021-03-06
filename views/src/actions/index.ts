import axios from 'axios'
import { ThunkActionSync } from '~/types/actions'
import Firebase from '~/libs/Firebase'

export const googleLogin = (isSignedIn: boolean): ThunkActionSync => async dispatch => {
    const { auth } = Firebase
    const provider = new auth.GoogleAuthProvider()
    provider.addScope('https://www.googleapis.com/auth/contacts.readonly')

    try {
        const { user } = await auth().signInWithPopup(provider)
        const token = await auth().currentUser?.getIdToken()

        if (token) {
            const photoUrl = user?.photoURL || ''
            const res = await axios.post(
                'http://localhost:8088/accounts/signup',
                { uid: user?.uid, emailSignUp: false, photoUrl },
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

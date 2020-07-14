import { Dispatch } from 'redux'

export type Action<P = {}> = {
    type: string
    payload?: P
}

export type ThunkAction = (dispatch: Dispatch) => void
export type ThunkActionSync = (dispatch: Dispatch) => Promise<void>

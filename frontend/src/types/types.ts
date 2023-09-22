import { User } from 'src/gql/types'

export type UserField<T> = {
  oldValue: T
  value: T
}

export type UserDataKey = Omit<User,'__typename' | 'id'>

export type UserDataFields = Record<keyof UserDataKey, UserField<User[keyof UserDataKey]>>;

export type UserData = {
  id: string
  fields: UserDataFields
  isNew: boolean
  isDelete: boolean
  selected: boolean
}

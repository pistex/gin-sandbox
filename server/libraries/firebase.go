package libraries

import (
	"mime/multipart"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

// FirebaseApp initialize firebase by credential.json.
func FirebaseApp() *firebase.App {
	panic("to be deleted")
}

// FirestoreDB return firestore client and error
func FirestoreDB() *firestore.Client {
	panic("to be deleted")
}

// FirestoreFind by collection and document ID.
func FirestoreFind(collecttion string, id string) (*firestore.DocumentSnapshot, error) {
	panic("to be deleted")
}

// FirestoreDelete by collection and document ID.
func FirestoreDelete(collecttion string, id string) (*firestore.WriteResult, error) {
	panic("to be deleted")
}

// FirestoreSearch by collection and condition
func FirestoreSearch(collecttion string, field string, condition string, property interface{}) ([]*firestore.DocumentSnapshot, error) {
	panic("to be deleted")
}

// FirestoreCreateOrSet by collection, id.
func FirestoreCreateOrSet(collecttion string, id string, data interface{}) (*firestore.WriteResult, error) {
	panic("to be deleted")
}

// FirestoreAdd by collection and automatically create id.
func FirestoreAdd(collecttion string, data interface{}) (*firestore.DocumentRef, *firestore.WriteResult, error) {
	panic("to be deleted")
}

// FirestoreUpdateField by collection, id, and field.
func FirestoreUpdateField(collecttion string, id string, field string, property interface{}) (*firestore.WriteResult, error) {
	panic("to be deleted")

}

// FirestoreUpdateFieldIfNotBlank by collection, id, and field.
func FirestoreUpdateFieldIfNotBlank(collecttion string, id string, field string, property interface{}) (*firestore.WriteResult, error) {
	panic("to be deleted")

}

// FirestoreDeleteField by collection, id, and field.
func FirestoreDeleteField(collecttion string, id string, field string) (*firestore.WriteResult, error) {
	panic("to be deleted")

}

// CloudStorageUpload to
func CloudStorageUpload(file multipart.File, path string) {
	panic("to be deleted")

}

func CreateProfilePicture(username string) {
	panic("to be deleted")

}

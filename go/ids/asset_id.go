package ids

type AssetID interface {
	ID
	IsAssetID()
}

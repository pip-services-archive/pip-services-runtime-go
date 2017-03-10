package runtime

type IComponent interface {
    Init(refs *References) error
    Open() error
    Close() error
}

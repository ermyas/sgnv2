package types

import (
	"gopkg.in/yaml.v2"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	MaxMonikerLength  = 70
	MaxIdentityLength = 3000
	MaxWebsiteLength  = 140
	MaxContactLength  = 140
	MaxDetailsLength  = 280

	// constant used in flags to indicate that description field should not be updated
	DoNotModifyDesc = "[do-not-modify]"
)

func NewDescription(moniker, identity, website, contact, details string) *Description {
	return &Description{
		Moniker:  moniker,
		Identity: identity,
		Website:  website,
		Contact:  contact,
		Details:  details,
	}
}

// String implements the Stringer interface for a Description object.
func (d *Description) String() string {
	out, _ := yaml.Marshal(d)
	return string(out)
}

// UpdateDescription updates the fields of a given description. An error is
// returned if the resulting description contains an invalid length.
func (d *Description) UpdateDescription(d2 *Description) error {
	err := d2.EnsureLength()
	if err != nil {
		return err
	}
	if d2.Moniker != DoNotModifyDesc {
		d.Moniker = d2.Moniker
	}
	if d2.Identity != DoNotModifyDesc {
		d.Identity = d2.Identity
	}
	if d2.Website != DoNotModifyDesc {
		d.Website = d2.Website
	}
	if d2.Contact != DoNotModifyDesc {
		d.Contact = d2.Contact
	}
	if d2.Details != DoNotModifyDesc {
		d.Details = d2.Details
	}
	return nil
}

// EnsureLength ensures the length of a validator's description.
func (d *Description) EnsureLength() error {
	if len(d.Moniker) > MaxMonikerLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid moniker length; got: %d, max: %d", len(d.Moniker), MaxMonikerLength)
	}

	if len(d.Identity) > MaxIdentityLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid identity length; got: %d, max: %d", len(d.Identity), MaxIdentityLength)
	}

	if len(d.Website) > MaxWebsiteLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid website length; got: %d, max: %d", len(d.Website), MaxWebsiteLength)
	}

	if len(d.Contact) > MaxContactLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid security contact length; got: %d, max: %d", len(d.Contact), MaxContactLength)
	}

	if len(d.Details) > MaxDetailsLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid details length; got: %d, max: %d", len(d.Details), MaxDetailsLength)
	}

	return nil
}

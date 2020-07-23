package fullcontact

type ProfileOptions func(*Profile)

type Profile struct {
	URL      string `json:"url,omitempty"`
	Username string `json:"username,omitempty"`
	UserId   string `json:"userid,omitempty"`
	Service  string `json:"service,omitempty"`
}

func NewProfile(options ...ProfileOptions) (*Profile, error) {
	profile := &Profile{}

	for _, opts := range options {
		opts(profile)
	}

	err := validateProfile(profile)
	if err != nil {
		profile = nil
	}
	return profile, err
}

func WithUrl(url string) ProfileOptions {
	return func(profile *Profile) {
		profile.URL = url
	}
}

func WithUsername(username string) ProfileOptions {
	return func(profile *Profile) {
		profile.Username = username
	}
}

func WithUserid(userid string) ProfileOptions {
	return func(profile *Profile) {
		profile.UserId = userid
	}
}

func WithService(service string) ProfileOptions {
	return func(profile *Profile) {
		profile.Service = service
	}
}

func validateProfile(profile *Profile) error {
	if isPopulated(profile.URL) {
		if isPopulated(profile.Username) || isPopulated(profile.UserId) {
			return NewFullContactError("Specifying username or userid together with url is not allowed")
		}
		return nil
	} else if isPopulated(profile.Service) {
		if isPopulated(profile.UserId) && isPopulated(profile.Username) {
			return NewFullContactError("Specifying userid together with username is not allowed")
		} else if isPopulated(profile.UserId) || isPopulated(profile.Username) {
			return nil
		} else {
			return NewFullContactError(
				"Either url or service plus username or userid must be set on every profiles entry.")
		}
	} else {
		return NewFullContactError(
			"Either url or service plus username or userid must be set on every profiles entry.")
	}
}

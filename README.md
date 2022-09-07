# ssi-sdk-mobile

This repository wraps the [ssi-sdk](https://github.com/TBD54566975/ssi-sdk) and adapts it to [gomobile's constraints](https://pkg.go.dev/golang.org/x/mobile/cmd/gobind#hdr-Type_restrictions) to produce Android and iOS binaries.

See [ssi-sdk/#181](https://github.com/TBD54566975/ssi-sdk/issues/181) for more information on why this wrapper is needed.

## Installation

You need [Android NDK](https://developer.android.com/ndk) for Android and xcodebuild for iOS
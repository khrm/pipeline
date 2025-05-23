// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Imports or reimports key material into an existing KMS key that was created
// without key material. ImportKeyMaterial also sets the expiration model and
// expiration date of the imported key material.
//
// By default, KMS keys are created with key material that KMS generates. This
// operation supports [Importing key material], an advanced feature that lets you generate and import the
// cryptographic key material for a KMS key. For more information about importing
// key material into KMS, see [Importing key material]in the Key Management Service Developer Guide.
//
// After you successfully import key material into a KMS key, you can [reimport the same key material] into that
// KMS key, but you cannot import different key material. You might reimport key
// material to replace key material that expired or key material that you deleted.
// You might also reimport key material to change the expiration model or
// expiration date of the key material.
//
// Each time you import key material into KMS, you can determine whether (
// ExpirationModel ) and when ( ValidTo ) the key material expires. To change the
// expiration of your key material, you must import it again, either by calling
// ImportKeyMaterial or using the import features of the KMS console.
//
// Before calling ImportKeyMaterial :
//
//   - Create or identify a KMS key with no key material. The KMS key must have an
//     Origin value of EXTERNAL , which indicates that the KMS key is designed for
//     imported key material.
//
// To create an new KMS key for imported key material, call the CreateKeyoperation with an
//
//	Origin value of EXTERNAL . You can create a symmetric encryption KMS key, HMAC
//	KMS key, asymmetric encryption KMS key, or asymmetric signing KMS key. You can
//	also import key material into a multi-Region keyof any supported type. However, you can't
//	import key material into a KMS key in a custom key store.
//
//	- Use the DescribeKeyoperation to verify that the KeyState of the KMS key is
//	PendingImport , which indicates that the KMS key has no key material.
//
// If you are reimporting the same key material into an existing KMS key, you
//
//	might need to call the DeleteImportedKeyMaterialto delete its existing key material.
//
//	- Call the GetParametersForImportoperation to get a public key and import token set for importing
//	key material.
//
//	- Use the public key in the GetParametersForImportresponse to encrypt your key material.
//
// Then, in an ImportKeyMaterial request, you submit your encrypted key material
// and import token. When calling this operation, you must specify the following
// values:
//
//   - The key ID or key ARN of the KMS key to associate with the imported key
//     material. Its Origin must be EXTERNAL and its KeyState must be PendingImport .
//     You cannot perform this operation on a KMS key in a custom key store, or on a KMS key in a
//     different Amazon Web Services account. To get the Origin and KeyState of a KMS
//     key, call DescribeKey.
//
//   - The encrypted key material.
//
//   - The import token that GetParametersForImportreturned. You must use a public key and token from
//     the same GetParametersForImport response.
//
//   - Whether the key material expires ( ExpirationModel ) and, if so, when (
//     ValidTo ). For help with this choice, see [Setting an expiration time]in the Key Management Service
//     Developer Guide.
//
// If you set an expiration date, KMS deletes the key material from the KMS key on
//
//	the specified date, making the KMS key unusable. To use the KMS key in
//	cryptographic operations again, you must reimport the same key material.
//	However, you can delete and reimport the key material at any time, including
//	before the key material expires. Each time you reimport, you can eliminate or
//	reset the expiration time.
//
// When this operation is successful, the key state of the KMS key changes from
// PendingImport to Enabled , and you can use the KMS key in cryptographic
// operations.
//
// If this operation fails, use the exception to help determine the problem. If
// the error is related to the key material, the import token, or wrapping key, use
// GetParametersForImportto get a new public key and import token for the KMS key and repeat the import
// procedure. For help, see [How To Import Key Material]in the Key Management Service Developer Guide.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see [Key states of KMS keys]in the Key Management Service Developer Guide.
//
// Cross-account use: No. You cannot perform this operation on a KMS key in a
// different Amazon Web Services account.
//
// Required permissions: [kms:ImportKeyMaterial] (key policy)
//
// Related operations:
//
// # DeleteImportedKeyMaterial
//
// # GetParametersForImport
//
// Eventual consistency: The KMS API follows an eventual consistency model. For
// more information, see [KMS eventual consistency].
//
// [Importing key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html
// [Key states of KMS keys]: https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html
// [How To Import Key Material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#importing-keys-overview
// [kms:ImportKeyMaterial]: https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html
// [reimport the same key material]: https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html#reimport-key-material
// [Setting an expiration time]: https://docs.aws.amazon.com/en_us/kms/latest/developerguide/importing-keys.html#importing-keys-expiration
// [KMS eventual consistency]: https://docs.aws.amazon.com/kms/latest/developerguide/programming-eventual-consistency.html
func (c *Client) ImportKeyMaterial(ctx context.Context, params *ImportKeyMaterialInput, optFns ...func(*Options)) (*ImportKeyMaterialOutput, error) {
	if params == nil {
		params = &ImportKeyMaterialInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportKeyMaterial", params, optFns, c.addOperationImportKeyMaterialMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportKeyMaterialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportKeyMaterialInput struct {

	// The encrypted key material to import. The key material must be encrypted under
	// the public wrapping key that GetParametersForImportreturned, using the wrapping algorithm that you
	// specified in the same GetParametersForImport request.
	//
	// This member is required.
	EncryptedKeyMaterial []byte

	// The import token that you received in the response to a previous GetParametersForImport request. It
	// must be from the same response that contained the public key that you used to
	// encrypt the key material.
	//
	// This member is required.
	ImportToken []byte

	// The identifier of the KMS key that will be associated with the imported key
	// material. This must be the same KMS key specified in the KeyID parameter of the
	// corresponding GetParametersForImportrequest. The Origin of the KMS key must be EXTERNAL and its
	// KeyState must be PendingImport .
	//
	// The KMS key can be a symmetric encryption KMS key, HMAC KMS key, asymmetric
	// encryption KMS key, or asymmetric signing KMS key, including a multi-Region keyof any supported
	// type. You cannot perform this operation on a KMS key in a custom key store, or
	// on a KMS key in a different Amazon Web Services account.
	//
	// Specify the key ID or key ARN of the KMS key.
	//
	// For example:
	//
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey.
	//
	// This member is required.
	KeyId *string

	// Specifies whether the key material expires. The default is KEY_MATERIAL_EXPIRES
	// . For help with this choice, see [Setting an expiration time]in the Key Management Service Developer Guide.
	//
	// When the value of ExpirationModel is KEY_MATERIAL_EXPIRES , you must specify a
	// value for the ValidTo parameter. When value is KEY_MATERIAL_DOES_NOT_EXPIRE ,
	// you must omit the ValidTo parameter.
	//
	// You cannot change the ExpirationModel or ValidTo values for the current import
	// after the request completes. To change either value, you must reimport the key
	// material.
	//
	// [Setting an expiration time]: https://docs.aws.amazon.com/en_us/kms/latest/developerguide/importing-keys.html#importing-keys-expiration
	ExpirationModel types.ExpirationModelType

	// The date and time when the imported key material expires. This parameter is
	// required when the value of the ExpirationModel parameter is KEY_MATERIAL_EXPIRES
	// . Otherwise it is not valid.
	//
	// The value of this parameter must be a future date and time. The maximum value
	// is 365 days from the request date.
	//
	// When the key material expires, KMS deletes the key material from the KMS key.
	// Without its key material, the KMS key is unusable. To use the KMS key in
	// cryptographic operations, you must reimport the same key material.
	//
	// You cannot change the ExpirationModel or ValidTo values for the current import
	// after the request completes. To change either value, you must delete (DeleteImportedKeyMaterial ) and
	// reimport the key material.
	ValidTo *time.Time

	noSmithyDocumentSerde
}

type ImportKeyMaterialOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportKeyMaterialMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportKeyMaterial{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportKeyMaterial{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportKeyMaterial"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpImportKeyMaterialValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportKeyMaterial(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opImportKeyMaterial(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportKeyMaterial",
	}
}

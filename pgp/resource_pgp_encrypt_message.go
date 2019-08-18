package gpg

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func resourcePGPEncryptMessage() *schema.Resource {
	return &schema.Resource{
		Create: resourceGPGEncryptedMessageCreate,
		Read:   resourceGPGEncryptedMessageRead,
		Delete: resourceGPGEncryptedMessageDelete,

		Schema: map[string]*schema.Schema{
			"content": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
				StateFunc: sha256sum,
			},
			"public_key": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
				StateFunc: sha256sum,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePGPEncryptMessage(d *schema.ResourceData, m interface{}) error {
	message := d.Get("content")
	public_key := d.Get("public_key")

	// Create public key entity
	publicKeyPacket, _ := pgp.GetPublicKeyPacket([]byte(public_key))
	pubEntity, _ := pgp.CreateEntityFromKeys(publicKeyPacket, nil)

	// Encrypt message
	encrypted, _ := pgp.Encrypt(pubEntity, []byte(message))

	return encrypted
}

func resourcePGPEncryptedMessageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceGPGEncryptedMessageDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

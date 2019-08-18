package pgp

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

func resourcePGPDecryptMessage() *schema.Resource {
	return &schema.Resource{
		Create: resourcePGPDecryptMessageCreate,
		Read:   resourcePGPDecryptMessageRead,
		Delete: resourcePGPDecryptMessageDelete,

		Schema: map[string]*schema.Schema{
			"content": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				ForceNew:  true,
				Sensitive: true,
				StateFunc: sha256sum,
			},
			"private_key": &schema.Schema{
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

func resourcePGPdecryptMessageCreate(d *schema.ResourceData, m interface{}) error {
	message := d.Get("content")
	public_key := d.Get("private_key")

	// Create private key entity
	privEntity, _ := pgp.GetEntity([]byte(TestPublicKey), []byte(TestPrivateKey))

	// Decrypt message
	decrypted, _ := pgp.Decrypt(privEntity, encrypted)

	return nil
}

func resourcePGPDecryptMessageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePGPDecryptMessageDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

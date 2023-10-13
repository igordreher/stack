<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


/**
 * BankAccountsCursor - OK
 * 
 * @package formance\stack\Models\Shared
 * @access public
 */
class BankAccountsCursor
{
	#[\JMS\Serializer\Annotation\SerializedName('cursor')]
    #[\JMS\Serializer\Annotation\Type('formance\stack\Models\Shared\BankAccountsCursorCursor')]
    public BankAccountsCursorCursor $cursor;
    
	public function __construct()
	{
		$this->cursor = new \formance\stack\Models\Shared\BankAccountsCursorCursor();
	}
}

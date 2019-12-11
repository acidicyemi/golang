<?php 
// A O(n^2) time and O(1) space program to find the longest palindromic substring 

// A utility function to print a substring str[low..high] 
function printSubStr($str, $low, $high) 
{ 
	for( $i = $low; $i <= $high; ++$i ) 
		echo $str[$i]; 
} 

// This function prints the longest palindrome substring (LPS) 
// of str[]. It also returns the length of the longest palindrome 
function longestPalSubstr($str) 
{ 
	$maxLength = 1; // The result (length of LPS) 

	$start = 0; 
	$len = strlen($str); 

	// One by one consider every character as center point of 
	// even and length palindromes 
	for ($i = 1; $i < $len; ++$i) 
	{ 
		// Find the longest even length palindrome with center points 
		// as i-1 and i. 
		$low = $i - 1; 
		$high = $i; 
		while ($low >= 0 && $high < $len && $str[$low] == $str[$high]) 
		{ 
			if ($high - $low + 1 > $maxLength) 
			{ 
				$start = $low; 
				$maxLength = $high - $low + 1; 
			} 
			--$low; 
			++$high; 
		} 

		// Find the longest odd length palindrome with center 
		// point as i 
		$low = $i - 1; 
		$high = $i + 1; 
		while ($low >= 0 && $high < $len && $str[$low] == $str[$high]) 
		{ 
			if ($high - $low + 1 > $maxLength) 
			{ 
				$start = $low; 
				$maxLength = $high - $low + 1; 
			} 
			--$low; 
			++$high; 
		} 
	} 

	echo "Longest palindrome substring is: "; 
	printSubStr($str, $start, $start + $maxLength - 1); 

	return $maxLength; 
} 

// Driver program to test above functions 

$str = "appearhhdjkdmdkddmdmdm"; 
echo "\nLength is: ". longestPalSubstr( $str ) ; 
return 0; 
?> 
forgeeksskeegforforgeeksskeegfor
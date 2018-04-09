#!/usr/bin/perl
use strict;
use warnings;

my $pattern = qr/^(\S+) +(\S+) +(\S+);(\S+) +(\S+) +(\S+)/;

my %translation;
while (<>) {
	chomp $_;
    if ( $_ =~ /$pattern/ ) {
        # push @{ $translation{$1} }, $4;
        # push @{ $translation{$2} }, $5;
        # push @{ $translation{$3} }, $6;
        $translation{$1}{$4}++;
        $translation{$2}{$5}++;
        $translation{$3}{$6}++;

    }
}

foreach my $spelling (sort keys %translation) {
	print "$spelling\n";
	foreach my $phone (sort keys %{$translation{$spelling}}) {
		print "    $phone: $translation{$spelling}{$phone}\n";
	}
}
#!/usr/bin/perl
use strict;
use warnings;

my $pattern = qr/^(\S+) +(\S+) +(\S+);(\S+) +(\S+) +(\S+)/;

my %translation;
while (<>) {
    chomp $_;
    my @spelling_or_phone = split( /;/, $_ );
    if ( scalar @spelling_or_phone == 2 ) {
        my ( @spelling, @phone );
        @spelling = split( / +/, $spelling_or_phone[0] );
        @phone    = split( / +/, $spelling_or_phone[1] );
        if ( scalar @spelling == scalar @phone ) {
            for ( my $t = 0 ; $t < scalar @spelling ; $t++ ) {
                if (! exists $translation{ $spelling[$t] }{ $phone[$t] }) {
                    $translation{ $spelling[$t] }{ $phone[$t] } = 0;
                }
                $translation{ $spelling[$t] }{ $phone[$t] }++;
               # print "$spelling[$t]\t$phone[$t]\n";
            }
            # print "\n";
        }
    }
}

foreach my $spelling ( sort keys %translation ) {
     print "$spelling\n";
    foreach my $phone ( sort keys %{ $translation{$spelling} } ) {
         print "    $phone: $translation{$spelling}{$phone}\n";
    }
}

#!/usr/bin/perl
use strict;
use warnings;

my $arbapet_pattern = qr/^(\S+)\t(\w+(?:[^\S\n]+\w+)*)$/msx;
my @line;
my %dict_list;
while (<>) {
    chomp $_;
    $dict_list{$_} = 0;
    push( @line, $_ );
}

# Repeat extraction until no more extractions can be made
my $old_lines = 0;
while ( scalar keys %dict_list > $old_lines ) {
    $old_lines = scalar keys %dict_list;
    print STDERR "$old_lines\n";    # debugging to track progress
    foreach my $line_to_find ( keys %dict_list ) {
        if ( $line_to_find =~ /$arbapet_pattern/msx ) {
            my ( $spelling, $pronounciation );
            ( $spelling, $pronounciation ) = ( $1, $2 );

            # my $pattern = qr/^$spelling(\S+)\t$pronounciation +(\w+(?: \w+)*)/;
            my $pattern = qr/^(\S+)$spelling\t(\w+(?: \w+)*) $pronounciation/;

            foreach my $line_to_find_in ( keys %dict_list ) {
                if ( $line_to_find_in =~ /$pattern/ ) {
                    $dict_list{"$1\t$2"} = 0;
                }
            }
        }
    }
    foreach ( sort keys %dict_list ) {
        print "$_\n";
    }
}

foreach ( sort keys %dict_list ) {
    print "$_\n";
}

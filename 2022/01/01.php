<?=max(array_map(fn($s)=>array_sum(explode("\n",$s)),explode("\n\n",file_get_contents('input.txt'))));

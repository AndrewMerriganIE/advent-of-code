package day.three;

import java.io.IOException;
import java.io.UnsupportedEncodingException;
import java.util.Arrays;
import java.util.Collections;

public class SpiralDistance {

	public static void main(final String[] args) throws UnsupportedEncodingException, IOException {
		
		final int target = 312051;
		int endValue = 0;
		int firstValue =0;
		
		for(int i=1; i<10000; i += 2){
			firstValue = endValue+  1;
			endValue = i*i;
			
			if(target <= endValue){
				break; 
			}	
		}
		
		final int dem = ((endValue - firstValue - 4) / 4) + 2;
		
		final int centerLines1 = endValue - dem/2;
		final int centerLines2 = centerLines1 - dem;
		final int centerLines3 = centerLines2 - dem;
		final int centerLines4 = centerLines3 - dem;
		
		final int distanceToaxis = Collections.min(
												Arrays.asList(
														Math.abs(centerLines1-target), 
														Math.abs(centerLines2-target), 
														Math.abs(centerLines3-target), 
														Math.abs(centerLines4-target)));
		
		final int distanceToCenter = distanceToaxis + dem/2;
		
		
		System.out.println(distanceToCenter);
		
	}

 
}

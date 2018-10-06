package day.three;

public class SpiralSum {

	private static final int[][] spiral = new int[101][101];

	public static void main(final String[] args) {

		final int target = 25;
		int dem = 1;
		int[] oldRing = new int[dem * dem];

		for (; dem < 10000; dem++) {

			final int[] newRing = new int[dem * dem];

			final int x = 0;
		
			
			final int value = sum(oldRing, newRing, index);
			newRing[x] = value;
			
			if (value > target) {
				break;
			}

			oldRing = newRing;
		}

		System.out.println(value);
	}

	private static int sum(final int indexX, final int indexY) {
		
		
		final int[] rowOne = spiral[indexX + 1];
		final int[] rowTwo = spiral[indexX];
		final int[] rowThree = spiral[indexX + 1];

		final int val1 = rowOne[indexY - 1] + rowOne[indexY] + rowOne[indexY + 1];
		final int val2 = rowTwo[indexY - 1] + rowTwo[indexY] + rowTwo[indexY + 1];
		final int val3 = rowThree[indexY - 1] + rowThree[indexY] + rowThree[indexY + 1];

		return val1 + val2 + val3;
	}

	

}
